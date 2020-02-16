import React from 'react';
import { Item, Button, Segment } from 'semantic-ui-react';
import { Link } from 'react-router-dom';
import { IRecipe } from '../../../app/models/recipe';

const RecipeListItem: React.FC<{ recipe: IRecipe }> = ({ recipe }) => {
  return (
    <Segment.Group>
      <Segment>
        <Item.Group>
          <Item>
            <Item.Image size='tiny' circular src='/assets/user.png' />
            <Item.Content>
              <Item.Header as='a'>{recipe.title}</Item.Header>
              <Item.Description>Hosted by Bob</Item.Description>
            </Item.Content>
          </Item>
        </Item.Group>
      </Segment>
      <Segment clearing>
        <span>{recipe.description}</span>
        <Button
          as={Link}
          to={`/recipe/${recipe.id}`}
          floated='right'
          content='View'
          color='blue'
        />
      </Segment>
    </Segment.Group>
  );
};

export default RecipeListItem;
