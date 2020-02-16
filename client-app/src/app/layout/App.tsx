import React, { Fragment, useContext, useEffect } from 'react';
import { ToastContainer } from 'react-toastify';
import LoginForm from '../../features/user/LoginForm';
import { Route, RouteComponentProps, Switch, withRouter } from 'react-router-dom';
import HomePage from '../../features/home/HomePage';
import NavBar from '../../features/nav/NavBar';
import { Container } from 'semantic-ui-react';
import RecipeForm from '../../features/recipes/form/RecipeForm';
import { RootStoreContext } from '../stores/rootStore';
import LoadingComponent from './LoadingComponent';
import ModalContainer from '../common/modals/ModalContainer';
import NotFound from './NotFound';
import RecipeDashboard from '../../features/recipes/dashboard/RecipeDashboard';
import { observer } from 'mobx-react-lite';
import RecipeDetails from '../../features/recipes/details/RecipeDetails';

const App: React.FC<RouteComponentProps> = ({ location }) => {
  const rootStore = useContext(RootStoreContext);
  const { setAppLoaded, appLoaded } = rootStore.commonStore;
  const { getUser } = rootStore.userStore;

  useEffect(() => {
    getUser()
    setAppLoaded()
  }, [getUser, setAppLoaded])

  if (!appLoaded) return <LoadingComponent content='Loading app...'/>

  return (
    <Fragment>
      <ModalContainer />
      <ToastContainer position='bottom-right' />
        <Route exact path="/" component={HomePage} />
        <Route
          path={"/(.+)"}
          render={() => (
            <Fragment>
                <NavBar />
                <Container style={{ marginTop: '7em' }}>
                  <Switch>
                    <Route exact path='/recipes' component={RecipeDashboard}/>
                    <Route path='/recipes/:id' component={RecipeDetails} />
                    <Route
                      key={location.key}
                      path={["/createRecipe", "/manage/:id"]}
                      component={RecipeForm}
                    />
                    <Route path='/login' component={LoginForm} />
                    <Route component={NotFound} />
                  </Switch>
                </Container>
            </Fragment>
          )}
        />
    </Fragment>
  );
}

export default withRouter(observer(App));
