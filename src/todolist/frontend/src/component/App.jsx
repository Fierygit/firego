import { TodoList } from './TodoList';
import { ThemeButton } from "./ThemeButton";
import './index.css';

function App() {

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <ThemeButton />
      <TodoList />
    </div>
  );
}

export default App;
