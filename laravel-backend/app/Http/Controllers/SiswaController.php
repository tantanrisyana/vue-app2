<?php

namespace App\Http\Controllers;

use App\Models\Siswa;
use Illuminate\Http\Request;

class SiswaController extends Controller
{
    public function index()
    {
        return Siswa::all();
    }

    public function store(Request $request)
    {
        $request->validate([
            'nama' => 'required|string',
            'umur' => 'required|integer',
            'alamat' => 'required|string',
        ]);

        return Siswa::create($request->all());
    }

    public function show(Siswa $siswa)
    {
        return $siswa;
    }

    public function update(Request $request, Siswa $siswa)
    {
        $request->validate([
            'nama' => 'required|string',
            'umur' => 'required|integer',
            'alamat' => 'required|string',
        ]);

        $siswa->update($request->all());
        return $siswa;
    }

    public function destroy(Siswa $siswa)
    {
        $siswa->delete();
        return response()->json(['message' => 'Siswa deleted']);
    }
}
