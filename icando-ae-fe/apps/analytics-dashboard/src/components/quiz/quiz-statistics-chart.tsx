import { PieChart } from "@mui/x-charts/PieChart";
import React from "react";

interface QuizStatisticsChartProps {
  pass: number;
  fail: number;
  isLoading: boolean;
}
export const QuizStatisticsChart: React.FC<QuizStatisticsChartProps> = ({
  pass,
  fail,
  isLoading,
}) => {
  return (
    <PieChart
      loading={isLoading}
      series={[
        {
          data: [
            {
              id: 0,
              value: pass,
              label: "Passed",
            },
            {
              id: 1,
              value: fail,
              label: "Failed",
            },
          ],
        },
      ]}
      width={600}
      height={300}
    />
  );
};
