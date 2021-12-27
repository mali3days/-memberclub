import React from 'react';

import { Loader } from '../Loader/Loader';

import './Table.css';

export function Table({ members, loading, className }) {
  function calcDate(value) {
    const date = new Date(value);
    return (
      date.getDate() +
      '.' +
      Number(date.getMonth() + 1) +
      '.' +
      date.getFullYear()
    );
  }

  return (
    <div style={{ width: '64%' }}>
      <h4>Members</h4>
      <div
        className="card"
        style={{
          overflow: 'auto',
          position: 'relative',
          maxHeight: 'calc(100vh - 260px)',
          minHeight: 357,
        }}
      >
        <Loader isActive={loading} />
        <table style={{ width: '100%' }}>
          <thead>
            <tr>
              <td>#</td>
              <td>Name</td>
              <td>Email</td>
              <td>Registration date</td>
            </tr>
          </thead>
          {(members.length > 0 || loading) ? (
            <tbody>
              {members.map((member) => {
                return (
                  <tr key={member.id}>
                    <td>{member.id}</td>
                    <td>{member.name}</td>
                    <td>{member.email}</td>
                    <td>{calcDate(member.registrationDate)}</td>
                  </tr>
                );
              })}
            </tbody>
          ) : (
            <tbody>
              <tr style={{ textAlign: 'center', width: '100%' }}>
                <td colSpan={4} style={{ color: '#ec5990', height: 224 }}>
                  No data to display
                </td>
              </tr>
              <tr style={{ textAlign: 'center', width: '100%' }}>
                <td
                  colSpan={4}
                  style={{
                    fontSize: 10,
                    position: 'absolute',
                    bottom: 0,
                    left: 0,
                  }}
                >
                  &#8592; Use "New member" form to create members
                </td>
              </tr>
            </tbody>
          )}
        </table>
      </div>
    </div>
  );
}
