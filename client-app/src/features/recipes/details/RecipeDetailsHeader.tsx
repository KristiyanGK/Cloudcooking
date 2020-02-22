import React from 'react';
import { Segment, Item, Header, Button, Image } from 'semantic-ui-react';
import { IRecipe } from '../../../app/models/recipe';
import { observer } from 'mobx-react-lite';
import { Link } from 'react-router-dom';
import { IUser } from '../../../app/models/user';

const recipeImageStyle = {
  filter: 'brightness(30%)'
};

const recipeImageTextStyle = {
  position: 'absolute',
  bottom: '5%',
  left: '5%',
  width: '100%',
  height: 'auto',
  color: 'white'
};

interface IProps {
  recipe: IRecipe;
  currUser: IUser | null;
}

const RecipeDetailedHeader: React.FC<IProps> = ({recipe, currUser}) => {

  return (
    <Segment.Group>
      <Segment basic attached='top' style={{ padding: '0' }}>
        <Image
          src={`/assets/categoryImages/${recipe.category}.jpg`}
          alt='temp'
          fluid
          style={recipeImageStyle}
        />
        <Segment style={recipeImageTextStyle} basic>
          <Item.Group>
            <Item>
              <Item.Content>
                <Header
                  size='huge'
                  content={recipe.title}
                  style={{ color: 'teal' }}
                />
                <p>
                  Made by <strong>{recipe.user}</strong>
                </p>
              </Item.Content>
            </Item>
          </Item.Group>
        </Segment>
      </Segment>
      <Segment clearing attached='bottom'>
        {currUser?.username === recipe.user && 
          <Button as={Link} to={`/manage/${recipe.id}`} color='orange' floated='right'>
            Manage Recipe
          </Button>
        }
      </Segment>
    </Segment.Group>
  );
};

export default observer(RecipeDetailedHeader);
