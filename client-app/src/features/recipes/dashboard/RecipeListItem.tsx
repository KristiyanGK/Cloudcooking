import React, { useContext } from 'react';
import { Item, Button, Segment } from 'semantic-ui-react';
import { Link } from 'react-router-dom';
import { IRecipe } from '../../../app/models/recipe';
import { RootStoreContext } from '../../../app/stores/rootStore';

const RecipeListItem: React.FC<{ recipe: IRecipe }> = ({ recipe }) => {
  const rootStore = useContext(RootStoreContext);

  const { user } = rootStore.userStore;
  const { deleteRecipe } = rootStore.recipeStore;

  return (
    <Segment.Group>
      <Segment>
        <Item.Group>
          <Item>
            <Item.Image size='tiny' circular src='/assets/user.png' />
            <Item.Content>
              <Item.Header as='a'>{recipe.title}</Item.Header>
                <Item.Description>Cooked by {recipe.user}</Item.Description>
                <Item.Extra>Category: {recipe.category.name}</Item.Extra>
            </Item.Content>
          </Item>
        </Item.Group>
      </Segment>
      <Segment clearing>
        <span>{recipe.description}</span>
        <Button
          as={Link}
          to={`/recipes/${recipe.id}`}
          floated='right'
          content='View'
          color='blue'
        />
        {user?.username == recipe.user && 
        <Button
          name={recipe.id}
          onClick={(e) => deleteRecipe(e, recipe.id)}
          floated='right'
          content='Delete'
          color='red'
        />}
      </Segment>
    </Segment.Group>
  );
};

export default RecipeListItem;
