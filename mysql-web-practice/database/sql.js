const mysql = require('mysql2');

const pool = mysql.createPool(
  process.env.JAWSDB_URL ?? {
    host: 'localhost',
    user: 'root',
    database: 'mydatabase',
    password: '1234',
    waitForConnections: true,
    connectionLimit: 10,
    queueLimit: 0
  }
);
const promisePool = pool.promise();

const sql = {

  getSections : async () => {
    const [rows] = await promisePool.query(`
      SELECT * FROM sections
    `)
    return rows
  },

}

module.exports = sql