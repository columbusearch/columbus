import React, { useState } from 'react'
import { MagnifyingGlassIcon } from '@heroicons/react/24/solid'
import Image from 'next/image'

const Home: React.FC = () => {
  const [query, setQuery] = useState('')
  const [results, setResults] = useState([])

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const res = await fetch(`/api?q=${query}`)
    const data = await res.json()
    setResults(data)
  }

  return (
    <div className="min-h-screen bg-gray-100">
      <div style={{ textAlign: "center", paddingTop: "2%" }}>
        <div style={{ justifyContent:"center", alignItems: "center", display: "flex" }}>
          <Image alt="Logo" src="/icon-go.svg" width={150} height={150} style={{marginRight:"1%"}}/>
          <h1 className="text-9xl font-bold text-gray-900">Columbus</h1>
        </div>
      </div>
      <div className="max-w-3xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <form onSubmit={handleSubmit} className="mt-6 flex">
          <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            className="flex-1 border border-gray-400 py-2 px-4 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500"
            placeholder="Search..."
          />
          <button type="submit" className="ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-cyan-400 hover:bg-cyan-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500">
            Search
          </button>
        </form>
        <div className="mt-8">
          {results.map((result: any) => (
            <div key={result.id} className="result p-4 bg-white shadow rounded-lg">
              <h2 className="text-lg font-medium text-gray-900">{result.title}</h2>
              <p className="mt-2 text-gray-600">{result.description}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default Home
