import { BarChart } from "@mui/x-charts/BarChart";
import { StudentCompetency } from "../../interfaces/competency";
import { useMemo } from "react";

interface CompetencyChartProps {
  data: StudentCompetency[];
}

export default function CompetencyChart({ data }: CompetencyChartProps) {
  const dataset = useMemo(() => {
    return data.map((each) => {
      return {
        ...each,
        incorrectCount: each.totalCount - each.correctCount,
      };
    });
  }, [data]);

  return (
    <BarChart
      margin={{ left: 100 }}
      dataset={dataset}
      yAxis={[{ scaleType: "band", dataKey: "competencyName" }]}
      series={[
        { dataKey: "correctCount", label: "Total Correct", stack: "A" },
        { dataKey: "incorrectCount", label: "Total Incorrect", stack: "A" },
      ]}
      layout="horizontal"
      width={600}
      height={400}
    />
  );
}
