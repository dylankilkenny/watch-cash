export default function validateToken(token) {
  return fetch('http://0.0.0.0:3001/api/validate', {
    method: 'POST',
    body: JSON.stringify({
      token: token
    }),
    headers: { 'Content-Type': 'application/json' }
  }).then(response => {
    return response.json();
  });
}
