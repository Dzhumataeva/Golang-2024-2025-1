// src/pages/TaskDetailPage.js Dzhumataeva Arukhan 
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';

const TaskDetailPage = () => {
    const { id } = useParams();
    const [task, setTask] = useState(null);

    useEffect(() => {
        axios.get(`http://localhost:8000/tasks/${id}`)
            .then(response => setTask(response.data))
            .catch(error => console.error('Error fetching task details!', error));
    }, [id]);

    return (
        <div>
            <h1>Task Details</h1>
            {task ? (
                <div>
                    <h2>{task.title}</h2>
                    <p>{task.description}</p>
                    <p><strong>Status:</strong> {task.status}</p>
                </div>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default TaskDetailPage;
