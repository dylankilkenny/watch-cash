import React from 'react';
import Table from 'react-bootstrap/Table';
import Button from 'react-bootstrap/Button';

const AddressTable = ({ subscribedAddresses, removeAddress }) => {
  return (
    <Table responsive>
      <thead>
        <tr>
          <th>#</th>
          <th>Address</th>
          <th>Date Added</th>
          <th />
        </tr>
      </thead>
      <tbody>
        {subscribedAddresses.map((item, i) => (
          <tr>
            <td>{i + 1}</td>
            <td>{item.address}</td>
            <td>{item.created_at}</td>
            <td>
              <Button
                onClick={() => removeAddress(item.address)}
                variant="danger"
                size="sm"
              >
                Remove
              </Button>
            </td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
};

export default AddressTable;
