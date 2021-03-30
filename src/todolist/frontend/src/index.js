import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import axios from "axios";

axios.interceptors.response.use((response) => {
  return response;
}, (error) => {
  console.log(error);
  if (401 === error?.response?.status) {
    window.location.href = error.response.data.redirect;
    return;
  }
  return error.response;
});

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
