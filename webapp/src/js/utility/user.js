export default {
  Validate: token => {
    return fetch('http://0.0.0.0:3001/api/validate', {
      method: 'POST',
      body: JSON.stringify({
        token: token
      }),
      headers: { 'Content-Type': 'application/json' }
    }).then(response => {
      return response.json();
    });
  },
  Login: (email, password) => {
    return fetch('http://0.0.0.0:3001/api/login', {
      method: 'POST',
      body: JSON.stringify({
        email: email,
        password: password
      }),
      headers: { 'Content-Type': 'application/json' }
    })
      .then(response => {
        return response.json();
      })
      .then(json => {
        if (json.status == 200) {
          sessionStorage.setItem('token', json.token);
          sessionStorage.setItem('firstname', json.firstname);
          return true;
        } else {
          return false;
        }
      });
  },
  Signup: (firstname, lastname, email, password) => {
    return fetch('http://0.0.0.0:3001/api/signup', {
      method: 'POST',
      body: JSON.stringify({
        firstname: firstname,
        lastname: lastname,
        email: email,
        password: password
      }),
      headers: { 'Content-Type': 'application/json' }
    })
      .then(response => {
        return response.json();
      })
      .then(json => {
        console.log(json);
        if (json.status == 200) {
          return true;
        } else {
          return json.code;
        }
      });
  },
  SubscribedAddresses: token => {
    const myHeaders = new Headers();
    myHeaders.append('Content-Type', 'application/json');
    myHeaders.append('Authorization', `Bearer ${token}`);
    return fetch('http://0.0.0.0:3001/api/private/address', {
      method: 'GET',
      credentials: 'include',
      headers: myHeaders
    })
      .then(response => {
        return response.json();
      })
      .then(json => {
        if (json.status == 200) {
          return json.addresses;
        } else {
          return json.code;
        }
      });
  }
};
