import React, { useContext, useEffect, useState } from 'react';
import { Grid, Loader } from 'semantic-ui-react';
import RecipeList from './RecipeList';
import { observer } from 'mobx-react-lite';
import LoadingComponent from '../../../app/layout/LoadingComponent';
import { RootStoreContext } from '../../../app/stores/rootStore';
import InfiniteScroll from 'react-infinite-scroller'
import RecipeFilters from './RecipeFilters';


const RecipeDashboard: React.FC = () => {

  const rootStore = useContext(RootStoreContext);
  const {loadRecipes, loadingInitial, setPage, page, totalPages} = rootStore.recipeStore;
  const {
    loadCategories
  } = rootStore.categoryStore;

  const [loadingNext, setLoadingNext] = useState(false);

  const handleGetNext = () => {
    setLoadingNext(true);
    setPage(page + 1);
    loadRecipes().then(() => setLoadingNext(false))
  }

  useEffect(() => {
    loadRecipes().then(() => {
      loadCategories()
    });
  }, [loadRecipes]);

  if (loadingInitial && page === 0)
    return <LoadingComponent content='Loading recipes' />;

  return (
    <Grid>
      <Grid.Column width={10}>
        <InfiniteScroll
            pageStart={0}
            loadMore={handleGetNext}
            hasMore={!loadingNext && page + 1 < totalPages}
            initialLoad={false}>
          <RecipeList />
        </InfiniteScroll>
      </Grid.Column>
      <Grid.Column width={6}>
        <RecipeFilters/>
      </Grid.Column>
      <Grid.Column width={10}>
        <Loader active={loadingNext}/>
      </Grid.Column>
    </Grid>
  );
};

export default observer(RecipeDashboard);
