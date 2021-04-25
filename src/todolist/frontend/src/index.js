import React from 'react';
import ReactDOM from 'react-dom';
import App from './component/App';
import axios from "axios";
import './index.css';

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

// tailwind css dark mode
if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.querySelector('html').classList.add('dark')
  localStorage.theme = 'dark';
} else {
  document.querySelector('html').classList.remove('dark')
  localStorage.theme = 'light';
}

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
