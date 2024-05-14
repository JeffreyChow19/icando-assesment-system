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
        normalizedCorrectScore: each.correctCount / each.totalCount,
        normalizedIncorrectScore:
          (each.totalCount - each.correctCount) / each.totalCount,
      };
    });
  }, [data]);

  return (
    <BarChart
      margin={{ left: 100 }}
      dataset={dataset}
      yAxis={[
        {
          scaleType: "band",
          dataKey: "competencyName",
        },
      ]}
      series={[
        {
          dataKey: "normalizedCorrectScore",
          label: "Passed Competencies",
          stack: "A",
          color: "#22c55e",
          valueFormatter: (passed) => `${(passed! * 100).toFixed(1)}%`,
        },
        {
          dataKey: "normalizedIncorrectScore",
          label: "Failed Competencies",
          stack: "A",
          color: "#e11d48",
          valueFormatter: (failed) => `${(failed! * 100).toFixed(1)}%`,
        },
      ]}
      layout="horizontal"
      width={500}
      height={400}
    />
  );
}
