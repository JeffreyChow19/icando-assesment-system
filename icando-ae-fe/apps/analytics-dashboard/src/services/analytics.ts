import {
  StudentData,
  StudentPerformance,
  StudentQuiz,
  StudentSubmissions,
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

export interface DashboardOverview{
  totalStudent: number;
  totalClass: number;
  totalOngoingQuiz: number;
}

export interface LatestSubmissions{
  data: StudentSubmissions[];
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

export const getOverview = async () => { 
  return (await api.get(`${path}/overview`)).data.data as DashboardOverview;
}

export const getLatestSubmissions = async () => {
  return (await api.get(`${path}/latest-submissions`)).data.data as StudentSubmissions[];
}