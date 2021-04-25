import { TodoList } from './TodoList';
import { ThemeButton } from "./ThemeButton";
import { DailyRecord } from './DailyRecord';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

function App() {

  return (
    <div className="min-h-screen bg-gradient-to-r from-gray-100 to-gray-300 dark:from-gray-700 dark:to-gray-900">
      <ThemeButton />
      <Router>
        <Switch>
          <Route path="/daily/:todo_id" exact>
            <DailyRecord />
          </Route>
          <Route path="/">
            <TodoList />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
