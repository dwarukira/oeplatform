import React, { FunctionComponent } from "react";
import { RouteComponentProps, navigate, Link } from "@reach/router";
import {
  PageHeader,
  Spin,
  Form,
  Input,
  Button,
  Checkbox,
  Row,
  Col,
  Select
} from "antd";
import { Title } from "../../componets/Heading";
import { useQuery } from "@apollo/react-hooks";
import { GET_USER, GET_ROLES } from "../../gql/queries";
import { Padding } from "../../componets/Padding";
import { Center } from "../../componets/Center";
import routes from "../../routes";

const { Option } = Select;

const UserEdit: FunctionComponent<RouteComponentProps> = (props: any) => {
  const { loading, data, error } = useQuery(GET_USER, {
    variables: {
      id: props.id
    }
  });

  const { loading: rloading, data: rdata, error: rerror } = useQuery(GET_ROLES);



  function onChange(value: any) {
    console.log(`selected ${value}`);
  }

  function onBlur() {
    console.log("blur");
  }

  function onFocus() {
    console.log("focus");
  }

  function onSearch(val: any) {
    console.log("search:", val);
  }

  const roles = () =>
  getRoles().map((role: any) => (
      <Option value={role.id}>{role.name}</Option>
    ));

  function getRoles() {
    if (rerror  || rloading) {
      return  [""]
    }

    return rdata.roles.list
  }

  function getUser() {
    if (error || loading) {
      return {
        roles: [""]
      };
    }
    return data.users.list[0];
  }

  return (
    <>
      <PageHeader
        onBack={() => navigate(routes.users_list)}
        title=""
        subTitle="Back to users accounts list"
      />
      {loading ? (
        <>
          <Spin />
        </>
      ) : (
        <>
          <Padding>
            {" "}
            <Title> Editing {getUser().name}'s User Account </Title>{" "}
          </Padding>
          <Padding>
            <Row gutter={8}>
              <Col span={8}>
                <Form layout="vertical">
                  <Form.Item label="Username">
                    <Input placeholder="Username" value={getUser().name} />
                  </Form.Item>
                  <Form.Item label="First name">
                    <Input
                      placeholder="First name"
                      value={getUser().firstName}
                    />
                  </Form.Item>

                  <Form.Item label="Last name">
                    <Input placeholder="Last Name" value={getUser().lastName} />
                  </Form.Item>

                  <Form.Item label="Email">
                    <Input placeholder="Email" value={getUser().email} />
                  </Form.Item>

                  <Form.Item label="Roles">
                    <Select
                      showSearch
                      defaultValue={getUser().roles.map(
                        (role: any) => role.name
                      )}
                      mode="tags"
                      style={{ width: "100%" }}
                      placeholder="Roles"
                      optionFilterProp="children"
                      onChange={onChange}
                      onFocus={onFocus}
                      onBlur={onBlur}
                      onSearch={onSearch}
                      filterOption={(input: any, option: any) =>
                        option.props.children
                          .toLowerCase()
                          .indexOf(input.toLowerCase()) >= 0
                      }
                    >
                      {roles()}
                    </Select>
                  </Form.Item>

                  <Form.Item label="">
                    <Checkbox checked={false}>
                      Send new password instructions
                    </Checkbox>
                  </Form.Item>

                  <Form.Item>
                    <Button type="primary">Save Changes</Button>
                  </Form.Item>
                </Form>
              </Col>
            </Row>
          </Padding>
        </>
      )}
    </>
  );
};

export default UserEdit;
