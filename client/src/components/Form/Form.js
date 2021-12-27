import React from 'react';
import { useForm } from 'react-hook-form';

import { Loader } from '../Loader/Loader';

import './Form.css';

export function Form({ onSubmit, loading, className }) {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  return (
    <div className="form-wrapper">
      <h4>New member</h4>
      <form onSubmit={handleSubmit(onSubmit)} className={className}>
        <label htmlFor="name">Name</label>
        <input id="name" {...register('name')} />

        <label htmlFor="email">Email</label>
        <input id="email" {...register('email')} />

        {errors.email && <p>This field is required</p>}

        <input type="submit" value="Add" />
        <input type="reset" value="Clear" />

        <Loader isActive={loading} />
      </form>
    </div>
  );
}
