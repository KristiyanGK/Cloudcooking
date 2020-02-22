import React, { useContext, useEffect } from 'react'
import { observer } from "mobx-react-lite";
import { RootStoreContext } from '../../../app/stores/rootStore';
import { RouteComponentProps } from 'react-router-dom';
import LoadingComponent from '../../../app/layout/LoadingComponent';
import { Grid } from 'semantic-ui-react';
import RecipeDetailsHeader from './RecipeDetailsHeader';
import RecipeDetailedInfo from './RecipeDetailedInfo';
import RecipeDetailedComments from './RecipeDetailedComments';

interface DetailParams {
    id: string;
}

const RecipeDetails: React.FC<RouteComponentProps<DetailParams>> = ({
    match,
    history
}) => {
    const rootStore = useContext(RootStoreContext);
    const { recipe, loadRecipe, loadingInitial } = rootStore.recipeStore;
    const { user } = rootStore.userStore

    useEffect(() => {
        loadRecipe(match.params.id)
    }, [loadRecipe, match.params.id, history]);

    if (loadingInitial) return <LoadingComponent content='Loading recipe...' />;

    if (!recipe) return <h2>Recipe not found!</h2>

    return (
        <Grid>
            <Grid.Column>
                <RecipeDetailsHeader recipe={recipe} currUser={user}/>
                <RecipeDetailedInfo recipe={recipe}/>
                <RecipeDetailedComments/>
            </Grid.Column>
        </Grid>
    )
}

export default observer(RecipeDetails);