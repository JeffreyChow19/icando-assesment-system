import React, {
  createContext,
  ReactElement,
  useContext,
  useEffect,
  useState,
} from "react";
import { useQuery } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { removeToken } from "../utils/local-storage.ts";
import { StudentQuiz } from "../interfaces/quiz.ts";
import { Student } from "../interfaces/user.ts";
import { getQuizAvailability } from "../services/quiz.ts";
import { getStudentProfile } from "../services/student.ts";

interface StudentQuizContextValue {
  studentQuiz: StudentQuiz | undefined;
  loading: boolean;
  setStudentQuiz: React.Dispatch<React.SetStateAction<StudentQuiz | undefined>>;
  refresh: () => void;
}
interface StudentProfileContextValue {
  student: Student | undefined;
  loading: boolean;
  setStudent: React.Dispatch<React.SetStateAction<Student | undefined>>;
  refresh: () => void;
}

export const StudentQuizContext = createContext<StudentQuizContextValue>({
  studentQuiz: undefined,
  loading: true,
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  setStudentQuiz: () => {},
  refresh: () => {},
});

export const StudentContext = createContext<StudentProfileContextValue>({
  student: undefined,
  loading: true,
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  setStudent: () => {},
  refresh: () => {},
});

export const QuizProvider = ({ children }: { children: ReactElement }) => {
  const [studentQuiz, setStudentQuiz] = useState<StudentQuiz>();
  const [loading, setLoading] = useState<boolean>(true);
  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["studentquiz"],
    queryFn: () => getQuizAvailability(),
    retry: false,
  });

  useEffect(() => {
    if (isLoading) {
      return;
    }

    setLoading(false);
    if (error && error instanceof AxiosError && error.response?.status == 401) {
      removeToken();
      setStudentQuiz(undefined);
      return;
    }

    if (!error && data) {
      setStudentQuiz(data);
    }
  }, [isLoading, data, error]);

  return (
    <StudentQuizContext.Provider
      value={{
        studentQuiz: studentQuiz,
        loading,
        setStudentQuiz: setStudentQuiz,
        refresh: () => refetch(),
      }}
    >
      {children}
    </StudentQuizContext.Provider>
  );
};

export const StudentProvider = ({ children }: { children: ReactElement }) => {
  const [student, setStudent] = useState<Student>();
  const [loading, setLoading] = useState<boolean>(true);
  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["student"],
    queryFn: () => getStudentProfile(),
    retry: false,
  });

  useEffect(() => {
    if (isLoading) {
      return;
    }

    setLoading(false);
    if (error && error instanceof AxiosError && error.response?.status == 401) {
      removeToken();
      setStudent(undefined);
      return;
    }

    if (!error && data) {
      setStudent(data);
    }
  }, [isLoading, data, error]);

  return (
    <StudentContext.Provider
      value={{
        student: student,
        loading,
        setStudent: setStudent,
        refresh: () => refetch(),
      }}
    >
      {children}
    </StudentContext.Provider>
  );
};

export const useStudentQuiz = () => useContext(StudentQuizContext);
export const useStudentProfile = () => useContext(StudentContext);
