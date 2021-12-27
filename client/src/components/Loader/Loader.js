import React from 'react';
import './Loader.css';

export function Loader({ isActive }) {
  if (!isActive) return null;

  return (
    <div className="loader">
      <div className="lds-dual-ring" />
    </div>
  );
}
