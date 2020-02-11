import React, { Fragment } from 'react';
import { ToastContainer } from 'react-toastify';
import LoginForm from '../../features/user/LoginForm';
import { Route, RouteComponentProps, Switch } from 'react-router-dom';
import HomePage from '../../features/home/HomePage';
import NavBar from '../../features/nav/NavBar';
import { Container } from 'semantic-ui-react';

const App: React.FC<RouteComponentProps> = ({ location }) => {
  return (
    <Fragment>
      <ToastContainer position='bottom-right'>
        <NavBar/>
        <Container style={{ margin: "7em" }}>
          <Route exact path="/" component={HomePage} />
            <Route
              path={"/(.+)"}
              render={() => (
                <Fragment>
                    <Switch>
                      <Route></Route>
                    </Switch>
                </Fragment>
              )}
            />
          <Route path='/login' component={LoginForm}/>
        </Container>
      </ToastContainer>
    </Fragment>
  );
}

export default App;
