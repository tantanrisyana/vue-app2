<?php

// app/Http/Controllers/SiswaApiController.php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Siswa;

class SiswaApiController extends Controller
{
    public function index()
    {
        return Siswa::all();
    }

    public function store(Request $request)
    {
        $request->validate([
            'nama' => 'required',
            'alamat' => 'required',
            'email' => 'required|email|unique:siswas',
        ]);

        $siswa = Siswa::create($request->all());

        return response()->json(['message' => 'Siswa created successfully', 'data' => $siswa], 201);
    }

    public function show(Siswa $siswa)
    {
        return $siswa;
    }

    public function update(Request $request, Siswa $siswa)
    {
        $request->validate([
            'nama' => 'required',
            'alamat' => 'required',
            'email' => 'required|email|unique:siswas,email,' . $siswa->id,
        ]);

        $siswa->update($request->all());

        return $siswa;
    }

    public function destroy(Siswa $siswa)
    {
        $siswa->delete();

        return response()->json(['message' => 'Siswa deleted successfully']);
    }
}
