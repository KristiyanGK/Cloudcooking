import React, { Fragment, useContext, useEffect } from 'react';
import { ToastContainer } from 'react-toastify';
import LoginForm from '../../features/user/LoginForm';
import { Route, RouteComponentProps, Switch } from 'react-router-dom';
import HomePage from '../../features/home/HomePage';
import NavBar from '../../features/nav/NavBar';
import { Container } from 'semantic-ui-react';
import RecipeForm from '../../features/recipes/form/RecipeForm';
import { RootStoreContext } from '../stores/rootStore';
import LoadingComponent from './LoadingComponent';
import ModalContainer from '../common/modals/ModalContainer';

const App: React.FC<RouteComponentProps> = ({ location }) => {
  const rootStore = useContext(RootStoreContext);
  const { setAppLoaded, token, appLoaded } = rootStore.commonStore;
  const { getUser } = rootStore.userStore;

  useEffect(() => {
    if (token) {
      getUser().finally(() => setAppLoaded)
    } else {
      setAppLoaded()
    }
  }, [getUser, setAppLoaded, token])

  if (!appLoaded) return <LoadingComponent content='Loading app...'/>

  return (
    <Fragment>
      <ModalContainer />
      <ToastContainer position='bottom-right'>
        <NavBar/>
        <Container style={{ margin: "7em" }}>
          <Route exact path="/" component={HomePage} />
            <Route
              path={"/(.+)"}
              render={() => (
                <Fragment>
                    <Switch>
                    <Route
                      key={location.key}
                      path={["/createRecipe", "/manage/:id"]}
                      component={RecipeForm}
                    />
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
