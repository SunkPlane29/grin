import * as React from 'react';
import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { View, Text, Button } from 'react-native';
import { RootStackParamList } from './RootStackParams';

type homeScreenProp = NativeStackNavigationProp<RootStackParamList, 'Login'>;

export default function HomeScreen() {
    const navigation = useNavigation<homeScreenProp>();

    return (
        <View style={{flex: 1, alignItems: 'center', justifyContent: 'center'}}>
            <Text>Home Screen</Text>
            <Button title="Login" onPress={() => navigation.navigate('Login')}></Button>
        </View>
    );
};