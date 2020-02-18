import React, { useContext, Fragment } from 'react';
import { Item, Label } from 'semantic-ui-react';
import { observer } from 'mobx-react-lite';
import RecipeListItem from './RecipeListItem';
import { RootStoreContext } from '../../../app/stores/rootStore';

const RecipeList: React.FC = () => {
  const rootStore = useContext(RootStoreContext);
  const { recipesByCategory } = rootStore.recipeStore;
  return (
    <Fragment>
      {recipesByCategory.map(([group, recipes]) => (
        <Fragment key={group}>
            <Label size='large' color='blue'>
                {group}
            </Label>
            <Item.Group divided>
                {recipes.map(recipe => (
                <RecipeListItem key={recipe.id} recipe={recipe} />
                ))}
            </Item.Group>
        </Fragment>
      ))}
    </Fragment>
  );
};

export default observer(RecipeList);
