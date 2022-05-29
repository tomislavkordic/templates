const express = require('express')
const app = express()
const port = 3000

const SLEEP_DURATION_SECONDS = 30

app.get('/', (req, res) => {
  res.send('Hello, World!')
})

app.get('/exit', (req, res) => {
  process.exit(1)
})

app.get('/sleep', async (req, res) => {
  await new Promise(resolve => setTimeout(resolve, SLEEP_DURATION_SECONDS * 1000));
  res.send(`Slept for ${SLEEP_DURATION_SECONDS} seconds.`)
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
