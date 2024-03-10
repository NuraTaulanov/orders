document.getElementById('orderForm').addEventListener('submit', function(e) {
    e.preventDefault(); // Предотвращение стандартного поведения формы

    const customerName = document.getElementById('customerName').value;
    const product = document.getElementById('product').value;

    // Создание тела запроса
    const orderData = {
        customerName: customerName,
        products: [
            { product_id: parseInt(product, 10) } // Пример, если продукт задается одним ID
        ]
    };

    fetch('http://localhost:8080/v1/order', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(orderData),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            document.getElementById('result').textContent = 'Order created successfully!';
        })
        .catch(error => {
            console.error('There has been a problem with your fetch operation:', error);
            document.getElementById('result').textContent = 'Failed to create order.';
        });
});
