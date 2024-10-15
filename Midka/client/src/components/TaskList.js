// src/components/TaskList.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import { Button, Card, Spinner, Alert, Modal, Row, Col } from 'react-bootstrap';

const TaskList = () => {
    const [tasks, setTasks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [taskToDelete, setTaskToDelete] = useState(null);  // Состояние для задачи, которую удаляем

    useEffect(() => {
        fetchTasks();
    }, []);

    const fetchTasks = () => {
        axios.get('http://localhost:8000/tasks')
            .then(response => {
                setTasks(response.data);
                setLoading(false);
            })
            .catch(error => {
                setError('Error fetching tasks');
                setLoading(false);
            });
    };

    const handleDelete = (id) => {
        axios.delete(`http://localhost:8000/tasks/${id}`)
            .then(() => {
                setTasks(tasks.filter(task => task.ID !== id));  // Обновляем список задач
                setTaskToDelete(null);  // Закрываем модальное окно
            })
            .catch(() => setError('Error deleting task'));
    };

    return (
        <div className="container">
            {/* Кнопка Add New Task сверху */}
            <div className="text-center mb-4">
                <Link to="/add" className="btn btn-primary shadow-sm">
                    <i className="fas fa-plus"></i> Add New Task
                </Link>
            </div>

            {/* Спиннер загрузки */}
            {loading && <Spinner animation="border" role="status"><span className="sr-only">Loading...</span></Spinner>}
            {error && <Alert variant="danger">{error}</Alert>}

            {/* Отображение списка задач */}  
            <Row className="gy-4">
                {tasks.map(task => (
                    <Col key={task.ID} md={6} lg={4}>
                        <Card className="shadow-sm h-100">
                            <Card.Body>
                                <Card.Title className="text-primary">{task.title}</Card.Title>
                                <Card.Text>
                                    {task.description.length > 100 
                                        ? task.description.substring(0, 100) + '...' 
                                        : task.description}
                                    <br />
                                    <span className="badge bg-info text-white">{task.status}</span>
                                </Card.Text>
                                <div className="d-flex justify-content-between mt-4">
                                    <Link to={`/edit/${task.ID}`} className="btn btn-outline-primary btn-sm">
                                        <i className="fas fa-edit"></i> Edit
                                    </Link>
                                    <Button variant="outline-danger" size="sm" onClick={() => setTaskToDelete(task)}>
                                        <i className="fas fa-trash"></i> Delete
                                    </Button>
                                    <Link to={`/tasks/${task.ID}`} className="btn btn-outline-info btn-sm">
                                        <i className="fas fa-info-circle"></i> Details
                                    </Link>
                                </div>
                            </Card.Body>
                        </Card>
                    </Col>
                ))}
            </Row>
            
            {/* Модальное окно для подтверждения удаления */}
            {taskToDelete && (
                <Modal show onHide={() => setTaskToDelete(null)} centered>
                    <Modal.Header closeButton>
                        <Modal.Title>Confirm Delete</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        Are you sure you want to delete the task <strong>{taskToDelete.title}</strong>?
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="secondary" onClick={() => setTaskToDelete(null)}>Cancel</Button>
                        <Button variant="danger" onClick={() => handleDelete(taskToDelete.ID)}>Delete</Button>
                    </Modal.Footer>
                </Modal>
            )}
        </div>
    );
};

export default TaskList;
