import React from 'react';
import { Alert } from 'react-bootstrap';

function ConditionalAlert(props) {
  if (props.isActive) {
    return (
      <Alert style={{ marginTop: 20 }} variant={props.hasStyle}>
        {props.content}
      </Alert>
    );
  } else {
    return <div style={{ marginTop: 20 }} />;
  }
}

export default ConditionalAlert;
