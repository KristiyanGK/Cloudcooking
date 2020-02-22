import React, { useContext, Fragment } from 'react';
import { Item, Label } from 'semantic-ui-react';
import { observer } from 'mobx-react-lite';
import RecipeListItem from './RecipeListItem';
import { RootStoreContext } from '../../../app/stores/rootStore';
import {format} from 'date-fns';

const RecipeList: React.FC = () => {
  const rootStore = useContext(RootStoreContext);
  const { recipesByDate } = rootStore.recipeStore;
  return (
    <Fragment>
      {recipesByDate.map(([group, recipes]) => (
        <Fragment key={group}>
            <Label size='large' color='blue'>
              {format(Date.parse(group), 'eeee do MMMM')}
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
