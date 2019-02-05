import React from 'react';
import Table from 'react-bootstrap/Table';

const AddressTable = ({ subscribedAddresses }) => {
  return (
    <Table responsive>
      <thead>
        <tr>
          <th>#</th>
          <th>Address</th>
          <th>Date Added</th>
        </tr>
      </thead>
      <tbody>
        {subscribedAddresses.map((item, i) => (
          <tr>
            <td>{i + 1}</td>
            <td>{item.address}</td>
            <td>{item.created_at}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
};

export default AddressTable;
