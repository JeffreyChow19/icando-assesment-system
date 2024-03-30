import { api } from "../utils/api.ts";
import { Meta } from "../interfaces/meta.ts";
import { Competency } from "../interfaces/competency.ts";

const path = "/designer/competency";

export interface GetAllCompetencyFilter {
  page?: number;
  limit?: number;
  query?: string;
}

export interface GetAllCompetencyResponse {
  meta: Meta;
  competencies: Competency[];
}

export const getAllCompetency = async (filter?: GetAllCompetencyFilter) => {
  return (await api.get(path, { params: filter }))
    .data as GetAllCompetencyResponse;
};
