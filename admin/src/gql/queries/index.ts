import gql from "graphql-tag";

export const GET_USERS = gql`
  query Users {
    users {
      count

      list {
        name
        email
        id
        updatedAt
        createdAt
        description
        lastLogin
        roles {
          id
          name
        }
      }
    }
  }
`;

export const GET_USER = gql`
  query User($id: ID) {
    users(id: $id) {
      count

      list {
        name
        email
        id
        updatedAt
        createdAt
        description
        lastLogin
        status
        roles {
          id
          name
        }
      }
    }
  }
`;

export const GET_ROLES = gql`
  query Roles {
    roles {
      list {
        description
        name
        createdAt
        updatedAt
        id

        permissions {
          id
          tag
          description
        }
      }
    }
  }
`;

export const GET_ROLE = gql`
  query Roles($id: ID) {
    roles(id: $id) {
      list {
        description
        name
        createdAt
        updatedAt
        id

        permissions {
          id
          tag
          description
        }
      }
    }
  }
`;

export const GET_SELLERS = gql`
  query Sellers($id: ID) {
    sellers(id: $id) {
      count
      list {
        id
        name
        createdAt
        updatedAt
        displayName
        phone

        user {
          name
          description
          createdAt
          updatedAt
          email
          id
        }
      }
    }
  }
`;
