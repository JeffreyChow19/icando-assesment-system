import {
  StudentData,
  StudentPerformance,
  StudentQuiz,
} from "../interfaces/student.ts";
import { api } from "../utils/api.ts";
import { StudentCompetency } from "../interfaces/competency.ts";

const path = "/teacher/analytics";
export interface PerformanceFilter {
  studentId?: string;
  quizId?: string;
}

export interface StudentStatisticsResponseData {
  studentInfo: StudentData;
  performance: StudentPerformance;
  competency: StudentCompetency[];
  quizzes: StudentQuiz[];
}

export const getStudentStatistics = async (
  studentId: string,
): Promise<StudentStatisticsResponseData> => {
  return (await api.get(`${path}/student/${studentId}`)).data
    .data as StudentStatisticsResponseData;
};

export const getPerformance = async (params: PerformanceFilter) => {
  return (await api.get(`${path}/performance`, { params })).data
    .data as StudentPerformance;
};
