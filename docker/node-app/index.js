const express = require("express");

const port = 8070;
const app = express();
app.listen(port, () => {
  console.log(`Server running at port:${port}/`);
});

app.get("/", (req, res, next) => {
  res.end("NODE-APP: served root path /");
});

app.get("/api/v2/content/path", (req, res, next) => {
  const name = req.query.name;
  res.end(`NODE-APP: /api/v2/content/path + name = ${name}`);
});

app.get("/api/v2/foo/bar", (req, res, next) => {
  const name = req.query.name;
  res.end("NODE-APP: /api/v2/foo/bar");
});
