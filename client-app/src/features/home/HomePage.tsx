import React, { useContext, Fragment } from 'react';
import { RootStoreContext } from '../../app/stores/rootStore';
import { Segment, Container, Header, Button, Image } from 'semantic-ui-react';
import { Link } from 'react-router-dom';
import LoginForm from '../user/LoginForm';
import RegisterForm from '../user/RegisterForm';
import { observer } from 'mobx-react-lite';

const HomePage = () => {
    const rootStore = useContext(RootStoreContext);
    const { user, isLoggedIn } = rootStore.userStore;
    const { openModal } = rootStore.modalStore;

    return (
        <Segment inverted textAlign='center' vertical className='masthead'>
            <Container text>
                <Header as='h1' inverted>
                <Image
                    size='massive'
                    src='/assets/logo.png'
                    alt='logo'
                    style={{ marginBottom: 12 }}
                />
                Cloudcooking
                </Header>
            {isLoggedIn && user ? (
                <Fragment>
                    <Header as='h2' inverted content={`Welcome back ${user.username}`}/>
                    <Button as={Link} to='/recipes' size='huge' inverted>
                        Go to recipes!
                    </Button>
                </Fragment>
            ) : (
                <Fragment>
                <Header as='h2' inverted content={`Welcome to Cloudcooking`} />
                <Button onClick={() => openModal(<LoginForm />)} size='huge' inverted>
                    Login
                </Button>
                <Button onClick={() => openModal(<RegisterForm />)} size='huge' inverted>
                    Register
                </Button>
                </Fragment>
            )}
            </Container>
        </Segment>
    );
};

export default observer(HomePage);