import { QuizCard } from '../../components/card/quiz-card.tsx';
import { Layout } from '../../layouts/layout.tsx';

export const Home = () => {
  return (
    <Layout pageTitle={'Home'} showTitle={false}>
      <h1 className="text-lg mb-2">Selamat datang, Andi</h1>
      <QuizCard quiz={ null} />
    </Layout>
  )
}
