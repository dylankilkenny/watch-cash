import React from 'react';
import { Alert } from 'react-bootstrap';

function ConditionalAlert(props) {
  if (props.isActive) {
    return <Alert bsStyle={props.hasStyle}>{props.content}</Alert>;
  } else {
    return null;
  }
}

export default ConditionalAlert;
