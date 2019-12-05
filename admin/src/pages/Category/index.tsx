import React, { FunctionComponent } from "react";

import { Table } from "antd";
import { RouteComponentProps, Link } from "@reach/router";
import { useQuery } from "@apollo/react-hooks";
import { GET_PRODUCTS, GET_CATEGORIES } from "../../gql/queries";
import { Title } from "../../componets/Heading";


function truncate(str: string) {
  return str.length > 10 ? str.substring(0, 7) + "..." : str;
}

const columns = [
  {
    title: "ID",
    dataIndex: "id",
    render: (text: string) => <Link to={`/product/${text}`}>{truncate(text)}</Link>
  },

  {
    title: "Name",
    dataIndex: "name"
  },
  {
    title: "Description",
    dataIndex: "description"
  },
  {
    title: "createdAt",
    dataIndex: "createdAt"
  },
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

const Category: FunctionComponent<RouteComponentProps> = () => {
  const { loading, error, data } = useQuery(GET_CATEGORIES);

  const hasData = data !== undefined && !error;
  function createColumnData() {
    return data.categories.list.map((item: any) => ({
      key: item.id,
      id: item.id,
      name: item.name,
      description: item.description,
      createdAt: new Date(item.createdAt).toLocaleString()
    }));
  }

  return (
    <>
      <Title> Listing all Category </Title>
      <Table
        loading={loading}
        rowSelection={rowSelection}
        columns={columns}
        dataSource={hasData ? createColumnData() : null}
      />
    </>
  );
};

export default Category;
