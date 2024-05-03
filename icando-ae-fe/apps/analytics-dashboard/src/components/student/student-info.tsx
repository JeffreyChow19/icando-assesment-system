import { StudentData } from "../../interfaces/student.ts";

interface StudentInfoProps {
  data: StudentData;
}

export const StudentInfo = ({ data }: StudentInfoProps) => {
  return (
    <div className="grid grid-cols-3 gap-y-2 gap-x-4 max-w-[400px] p-2 divide-y">
      <div className="text-right font-semibold">NISN</div>
      <div className="col-span-2">{data.student.nisn}</div>

      <div className="text-right font-semibold">Name</div>
      <div className="col-span-2">
        {data.student
          ? `${data.student.firstName} ${data.student.lastName}`
          : ""}
      </div>

      <div className="text-right font-semibold">Email</div>
      <div className="col-span-2">{data.student.email}</div>

      <div className="text-right font-semibold">Class</div>
      <div className="col-span-2">
        {data.student ? `${data.class.grade} - ${data.class.name}` : ""}
      </div>
    </div>
  );
};
