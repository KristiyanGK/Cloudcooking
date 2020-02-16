import React, { useContext, useEffect } from 'react';
import { Grid } from 'semantic-ui-react';
import RecipeList from './RecipeList';
import { observer } from 'mobx-react-lite';
import LoadingComponent from '../../../app/layout/LoadingComponent';
import { RootStoreContext } from '../../../app/stores/rootStore';

const RecipeDashboard: React.FC = () => {

  const rootStore = useContext(RootStoreContext);
  const {loadRecipes, loadingInitial} = rootStore.recipeStore;

  useEffect(() => {
    loadRecipes();
  }, [loadRecipes]);

  if (loadingInitial)
    return <LoadingComponent content='Loading recipes' />;

  return (
    <Grid>
      <Grid.Column width={10}>
        <RecipeList />
      </Grid.Column>
      <Grid.Column width={6}>
        <h2>Recipe filters</h2>
      </Grid.Column>
    </Grid>
  );
};

export default observer(RecipeDashboard);
