// src/pages/HomePage.js Dzhumataeva Arukhan 
import React from 'react';
import TaskList from '../components/TaskList';
import { Link } from 'react-router-dom';

const HomePage = () => {
    return (
        <div>
            <h1>Task List</h1>
            <TaskList />
        </div>
    );
};

export default HomePage;
