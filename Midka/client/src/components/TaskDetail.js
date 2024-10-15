// src/components/TaskDetail.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { Spinner, Alert, Card, Container } from 'react-bootstrap';

const TaskDetail = () => {
    const { id } = useParams(); // Получаем ID задачи из параметров URL
    const [task, setTask] = useState(null); // Состояние для хранения данных задачи
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

   useEffect(() => {
    axios.get(`http://localhost:8000/tasks/${id}`)
        .then(response => {
            setTask(response.data); // Сохраняем данные задачи в состоянии
            setLoading(false); // Отключаем индикатор загрузки
        })
        .catch(() => {
            setError('Error fetching task details'); // Если произошла ошибка, отображаем сообщение
            setLoading(false);
        });
    }, [id]); // Выполняется каждый раз при изменении ID задачи


    return (
        <Container className="mt-5">
            {/* Спиннер загрузки */}
            {loading && <Spinner animation="border" role="status"><span className="sr-only">Loading...</span></Spinner>}
            {error && <Alert variant="danger">{error}</Alert>}

            {/* Карточка с детальной информацией о задаче */}
            {task && (
                <Card className="shadow-sm p-4">
                    <Card.Body>
                        <Card.Title className="text-center mb-4 text-primary">{task.title}</Card.Title>
                        <Card.Text className="text-center">
                            <strong>Description: </strong>{task.description}
                        </Card.Text>
                        <Card.Text className="text-center">
                            <strong>Status: </strong><span className="badge bg-info text-white">{task.status}</span>
                        </Card.Text>
                    </Card.Body>
                </Card>
            )}
        </Container>
    );
};

export default TaskDetail;
