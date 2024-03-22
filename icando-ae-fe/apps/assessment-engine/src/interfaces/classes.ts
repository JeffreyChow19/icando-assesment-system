import { Student } from "./student";
import { Teacher } from "./teacher";

export interface Class {
  id: string;
  name: string;
  grade: string;
  institutionId: string;
  students?: Student[];
  teachers?: Teacher[];
}
