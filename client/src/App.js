import React from 'react';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import './App.css';

import { Form } from './components/Form/Form';
import { Table } from './components/Table/Table';

export function App() {
  const [members, setMembers] = React.useState([]);
  const [loadingForm, setLoadingForm] = React.useState(false);
  const [loadingTable, setLoadingTable] = React.useState(false);

  function getAllMembers({ withScroll } = {}) {
    setLoadingTable(true);
    (async () => {
      const data = await fetch('http://localhost:4000').then((res) =>
        res.json()
      );
      setLoadingTable(false);
      setMembers(data);

      if (withScroll) {
        document.querySelector('#content > div:nth-child(2) > div').scroll(0, Number.MAX_SAFE_INTEGER)
      }
    })();
  }

  function handleSubmit({ name, email }) {
    setLoadingForm(true);
    (async () => {
      try {
        const status = await fetch('http://localhost:4000', {
          method: 'POST',
          body: JSON.stringify({
            name: name.trim(),
            email: email.trim(),
          }),
        })
          .then(function (response) {
            if (!response.ok) {
              return response.text().then((text) => {
                toast.error(text, {
                  position: 'top-right',
                  autoClose: 5000,
                  hideProgressBar: true,
                  closeOnClick: true,
                  pauseOnHover: true,
                  draggable: false,
                  progress: false,
                });
                throw Error(text);
              });
            }
            return response.status;
          })
          .then(function (response) {
            toast.success('Member added successfully!', {
              position: 'top-right',
              autoClose: 5000,
              hideProgressBar: true,
              closeOnClick: true,
              pauseOnHover: true,
              draggable: false,
              progress: false,
            });
            return response;
          })
          .catch(function (error) {
            toast.error(error, {
              position: 'top-right',
              autoClose: 5000,
              hideProgressBar: true,
              closeOnClick: true,
              pauseOnHover: true,
              draggable: false,
              progress: false,
            });
          });

        if (status === 200) getAllMembers({ withScroll: true });
      } catch (err) {
        toast.error(err, {
          position: 'top-right',
          autoClose: 5000,
          hideProgressBar: true,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: false,
          progress: false,
        });
      }
      setLoadingForm(false);
    })();
  }

  React.useEffect(getAllMembers, []);

  return (
    <div className="App">
      <header>
        <h1>Welcome to the Club!</h1>
      </header>
      <main>
        <div id="content">
          <Form
            onSubmit={handleSubmit}
            loading={loadingForm}
            className="card"
          />
          <Table members={members} loading={loadingTable} className="card" />
        </div>
      </main>
      <ToastContainer />
    </div>
  );
}
