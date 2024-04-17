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
import { getQuizAvailability } from "../services/quiz.ts";

interface UserContextValue {
  studentQuiz: StudentQuiz | undefined;
  loading: boolean;
  setStudentQuiz: React.Dispatch<React.SetStateAction<StudentQuiz | undefined>>;
  refresh: () => void;
}

export const UserContext = createContext<UserContextValue>({
  studentQuiz: undefined,
  loading: true,
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  setStudentQuiz: () => {},
  refresh: () => {},
});

export const UserProvider = ({ children }: { children: ReactElement }) => {
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
    <UserContext.Provider
      value={{
        studentQuiz: studentQuiz,
        loading,
        setStudentQuiz: setStudentQuiz,
        refresh: () => refetch(),
      }}
    >
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => useContext(UserContext);
