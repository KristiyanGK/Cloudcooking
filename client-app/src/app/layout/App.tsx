import React, { Fragment } from 'react';
import { ToastContainer } from 'react-toastify';
import LoginForm from '../../features/user/LoginForm';
import { Route } from 'react-router-dom';

const App = () => {
  return (
    <Fragment>
      <ToastContainer position='bottom-right'>
        <h1>Hello there</h1>
        <Route path='/login' component={LoginForm}/>
      </ToastContainer>
    </Fragment>
  );
}

export default App;
