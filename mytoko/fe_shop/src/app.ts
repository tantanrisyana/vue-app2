import express, { Request, Response } from 'express';
import bodyParser from 'body-parser';
import axios from 'axios';
import mysql from 'mysql2/promise';

const app = express();
const PORT = 3000;

// Middleware
app.use(bodyParser.json());

// Database connection configuration
const dbConfig = {
  host: 'localhost',
  user: 'root',
  password: '',
  database: 'shops',
};

// Function to create a MySQL connection
const createConnection = async () => {
  const connection = await mysql.createConnection(dbConfig);
  return connection;
};

// Routes
// Users
app.get('/users', async (req: Request, res: Response) => {
  try {
    const connection = await createConnection();
    const [rows] = await connection.execute('SELECT * FROM users');
    connection.end();
    res.json(rows);
  } catch (error) {
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

app.post('/users', async (req: Request, res: Response) => {
  try {
    const newUser = req.body;
    const connection = await createConnection();
    await connection.execute('INSERT INTO users (name, email) VALUES (?, ?)', [
      newUser.name,
      newUser.email,
    ]);
    connection.end();
    res.json(newUser);
  } catch (error) {
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

// Products
app.get('/products', async (req: Request, res: Response) => {
  try {
    const connection = await createConnection();
    const [rows] = await connection.execute('SELECT * FROM products');
    connection.end();
    res.json(rows);
  } catch (error) {
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

app.post('/products', async (req: Request, res: Response) => {
  try {
    const newProduct = req.body;
    const connection = await createConnection();
    await connection.execute('INSERT INTO products (name, price) VALUES (?, ?)', [
      newProduct.name,
      newProduct.price,
    ]);
    connection.end();
    res.json(newProduct);
  } catch (error) {
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

// Start the server
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
