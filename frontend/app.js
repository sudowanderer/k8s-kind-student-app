const API_URL = '/api/v1/students';

document.addEventListener('DOMContentLoaded', () => {
    const studentForm = document.getElementById('studentForm');
    const studentList = document.getElementById('studentList');

    // Load students on page load
    fetchStudents();

    // Handle form submission
    studentForm.addEventListener('submit', async (event) => {
        event.preventDefault();

        const name = document.getElementById('name').value;
        const age = document.getElementById('age').value;

        if (name && age) {
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name, age: parseInt(age, 10) }),
            });

            if (response.ok) {
                fetchStudents();
                studentForm.reset();
            } else {
                alert('Failed to add student');
            }
        }
    });

    // Fetch and display students
    async function fetchStudents() {
        const response = await fetch(API_URL);
        if (response.ok) {
            const students = await response.json();
            renderStudents(students);
        } else {
            alert('Failed to load students');
        }
    }

    // Render students in the list
    function renderStudents(students) {
        studentList.innerHTML = '';
        students.forEach((student) => {
            const li = document.createElement('li');
            li.textContent = `${student.name} (${student.age} years old)`;
            studentList.appendChild(li);
        });
    }
});