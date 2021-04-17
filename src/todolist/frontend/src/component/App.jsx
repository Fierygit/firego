import { TodoList } from './TodoList';
import { ThemeButton } from "./ThemeButton";
import '../index.css';

function App() {

  return (
    <div className="min-h-screen relative bg-gradient-to-r from-gray-100 to-gray-300 dark:from-gray-700 dark:to-gray-900">
      <ThemeButton />
      <TodoList />
    </div>
  );
}

export default App;
