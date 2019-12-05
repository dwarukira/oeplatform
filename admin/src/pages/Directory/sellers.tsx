import React, { FunctionComponent } from "react";
import { RouteComponentProps, Link } from "@reach/router";

import { Table, Col, Row } from "antd";
import { renderCardContent } from "../../componets/Card";
import { useQuery } from "@apollo/react-hooks";
import { Title } from "../../componets/Heading";
import { GET_SELLERS } from "../../gql/queries";
import { PaddingTop } from "../../componets/Padding";

function truncate(str: string) {
  return str.length > 10 ? str.substring(0, 7) + "..." : str;
}

const columns = [
  {
    title: "ID",
    dataIndex: "id",
    render: (text: any) => (
      <Link to={`/directory/seller/${text}`}>{truncate(text)}</Link>
    )
  },
  {
    title: "Name",
    dataIndex: "name"
  },

  {
    title: "Display Name",
    dataIndex: "displayName"
  },

  {
    title: "Email",
    dataIndex: "email"
  },

  {
    title: "Phone",
    dataIndex: "phone"
  },
  {
    title: "Member Since",
    dataIndex: "createdAt"
  },

  {
    title: "Last modified on",
    dataIndex: "updateAt"
  }
];

// rowSelection object indicates the need for row selection
const rowSelection = {
  onChange: (selectedRowKeys: any, selectedRows: any) => {
    console.log(
      `selectedRowKeys: ${selectedRowKeys}`,
      "selectedRows: ",
      selectedRows
    );
  },
  getCheckboxProps: (record: { name: string }) => ({
    disabled: record.name === "Disabled User", // Column configuration not to be checked
    name: record.name
  })
};

const SellerList: FunctionComponent<RouteComponentProps> = (): JSX.Element => {
  const { loading, error, data } = useQuery(GET_SELLERS);

  const content = {
    title: "Total number of sellers",
    total: data ? data.sellers.count : 0
  };

  const hasData = data !== undefined;
  function createColumnData() {
    return data.sellers.list.map((item: any) => ({
      key: item.id,
      id: item.id,
      name: item.name,
      phone: item.phone,
      email: item.user.email,
      displayName: item.displayName,
      createdAt: new Date(item.createdAt).toLocaleString(),
      updateAt: new Date(item.updatedAt).toLocaleString()
    }));
  }

  if (error) {
    return <p> XX </p>;
  }

  return (
    <div>
      <Title> Listing all sellers </Title>
      <PaddingTop />
      <Row gutter={16}>
        <Col className="gutter-row" span={6}>
          {renderCardContent(content, "1")}
        </Col>
      </Row>

      <PaddingTop />
      <Table
        loading={loading}
        rowSelection={rowSelection}
        columns={columns}
        dataSource={hasData ? createColumnData() : null}
      />
    </div>
  );
};

export default SellerList;
