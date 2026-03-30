let token = localStorage.getItem('token');

document.getElementById('loginBtn').addEventListener('click', async () => {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const response = await fetch(`${apiUrl}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
    });
    if (response.ok) {
        const data = await response.json();
        token = data.token;
        localStorage.setItem('token', token);
        document.getElementById('auth').style.display = 'none';
        document.getElementById('app').style.display = 'block';
        loadBooks();
    } else {
        alert('Login fallido');
    }
});

