import { Layout } from "../layouts/layout.tsx";
import { CheckIcon } from "lucide-react";

export const Submit = () => {
  return (
    <Layout pageTitle="Quiz" showTitle={false} showNavigation={false}>
      <div className="flex flex-col items-center justify-center gap-1 text-center">
        <div className="my-6 text-white rounded-full bg-green-500 p-2">
          <CheckIcon className="w-8 h-8" />
        </div>
        <h1 className="font-bold text-2xl">Jawaban telah terkirim!</h1>
        <h2 className="font-semibold">
          Terima kasih telah mengerjakan quiz ini
        </h2>
      </div>
    </Layout>
  );
};
