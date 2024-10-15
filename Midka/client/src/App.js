// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import AddTaskPage from './pages/AddTaskPage';
import EditTaskPage from './pages/EditTaskPage';
import TaskDetailPage from './pages/TaskDetailPage';

const App = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/add" element={<AddTaskPage />} />
                <Route path="/edit/:id" element={<EditTaskPage />} />
                <Route path="/tasks/:id" element={<TaskDetailPage />} />
            </Routes>
        </Router>
    );
};

export default App;
