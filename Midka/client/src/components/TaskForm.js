// src/components/TaskForm.js
import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { Form, Button, Alert } from 'react-bootstrap';

const TaskForm = ({ existingTask }) => { // Используем существующую задачу, если редактируем Арухан
    const [title, setTitle] = useState(existingTask ? existingTask.title : ''); // Устанавливаем начальные значения
    const [description, setDescription] = useState(existingTask ? existingTask.description : '');
    const [status, setStatus] = useState(existingTask ? existingTask.status : 'pending');
    const [error, setError] = useState(null);
    const navigate = useNavigate(); // Хук для навигации

    const handleSubmit = (e) => {
        e.preventDefault();
        const task = { title, description, status };

        const request = existingTask
        ? axios.put(`http://localhost:8000/tasks/${existingTask.ID}`, task) //Если редактируем, то отправляем PUT запрос
        : axios.post('http://localhost:8000/tasks', task); // Если создаем новую задачу, то отправляем POST запрос

        request
            .then(() => navigate('/'))
            .catch(() => setError('Error saving task'));
    };

    return (
        <div className="container">
            <h2>{existingTask ? "Edit Task" : "Add New Task"}</h2>

            {error && <Alert variant="danger">{error}</Alert>}

            <Form onSubmit={handleSubmit}>
                <Form.Group className="mb-3">
                    <Form.Label>Title</Form.Label>
                    <Form.Control 
                        type="text" 
                        value={title} 
                        onChange={(e) => setTitle(e.target.value)} 
                        required 
                    />
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Description</Form.Label>
                    <Form.Control 
                        as="textarea" 
                        rows={3}
                        value={description} 
                        onChange={(e) => setDescription(e.target.value)} 
                        required 
                    />
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Status</Form.Label>
                    <Form.Select value={status} onChange={(e) => setStatus(e.target.value)}>
                        <option value="pending">Pending</option>
                        <option value="in-progress">In Progress</option>
                        <option value="completed">Completed</option>
                    </Form.Select>
                </Form.Group>
                <Button variant="primary" type="submit">
                {existingTask ? "Update Task" : "Create Task"} 
                 </Button>

            </Form>
     #   </div>
    );
};

export default TaskForm;
