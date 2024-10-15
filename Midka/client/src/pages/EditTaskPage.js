// src/pages/EditTaskPage.js
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import TaskForm from '../components/TaskForm';
import axios from 'axios';

const EditTaskPage = () => {
    const { id } = useParams();
    const [existingTask, setExistingTask] = useState(null);

    useEffect(() => {
        axios.get(`http://localhost:8000/tasks/${id}`)
            .then(response => setExistingTask(response.data))
            .catch(error => console.error('Error fetching task details!', error));
    }, [id]);

    return (
        <div>
            <h1>Edit Task</h1>
            {existingTask ? <TaskForm existingTask={existingTask} /> : <p>Loading...</p>}
        </div>
    );
};

export default EditTaskPage;
